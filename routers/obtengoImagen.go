package routers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/Hoeru23/twittergo/awsgo"
	"github.com/Hoeru23/twittergo/bd"
	"github.com/Hoeru23/twittergo/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func ObtenerImagen(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var r models.RestApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Usuario no encontrado > " + err.Error()
		return r
	}

	var filename string
	switch uploadType {
	case "A":
		filename = "avatars/" + perfil.Avatar
	case "B":
		filename = "banners/" + perfil.Banner
	}

	fmt.Println("Filename " + filename)
	svc := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFromS3(ctx, svc, filename)
	if err != nil {
		r.Status = 500
		r.Message = "Ocurrió un error al descargar el archivo de S3 > " + err.Error()
		return r
	}

	r.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       file.String(),
		Headers: map[string]string{
			"Content-Type":        "application/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
		},
	}

	return r
}

func downloadFromS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	defer obj.Body.Close()
	fmt.Println("Bucketname = " + bucket)

	file, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)
	return buffer, nil
}
