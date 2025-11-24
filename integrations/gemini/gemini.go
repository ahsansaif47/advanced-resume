package gemini

import (
	"context"
	"errors"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ahsansaif47/advanced-resume/config"
	"google.golang.org/genai"
)

func GenAIClient() (*genai.Client, error) {

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: config.GetConfig().GeminiAPIKey,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func sanitizeFileName(name string) string {
	name = strings.ToLower(name)

	re := regexp.MustCompile(`[^a-z0-9-]+`)
	name = re.ReplaceAllString(name, "-")
	name = strings.Trim(name, "-")
	return name

}

func GetResponse(c *genai.Client, path string) (string, error) {
	ctx := context.Background()

	mimeType := "image/png"

	file, err := c.Files.UploadFromPath(
		ctx,
		path,
		&genai.UploadFileConfig{
			Name:        sanitizeFileName(filepath.Base(path)),
			DisplayName: filepath.Base(path),
			MIMEType:    mimeType,
		},
	)

	if err != nil {
		var apiErr *genai.APIError = &genai.APIError{}
		if errors.As(err, apiErr) {
			if apiErr.Code == 409 {
				// file already exists → delete → re-upload
				c.Files.Delete(ctx, sanitizeFileName(filepath.Base(path)), nil)

				// retry upload
				file, err = c.Files.UploadFromPath(ctx, path, &genai.UploadFileConfig{
					Name:        sanitizeFileName(filepath.Base(path)),
					DisplayName: filepath.Base(path),
					MIMEType:    mimeType,
				})

				if err != nil {
					return "", err
				}

			}
		}
	}

	prompt := []*genai.Part{
		genai.NewPartFromURI(file.URI, file.MIMEType),
		// genai.NewPartFromText("List down the Job Titles, skills and locations."),
		genai.NewPartFromText("Parse the resume to extract the person details like personal information, work experience, skills and all other detials listed on the resume. Return the response structured as json"),
	}

	contents := genai.Content{
		Parts: prompt,
	}

	// Generate content WITH the file
	result, err := c.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		[]*genai.Content{&contents},
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
