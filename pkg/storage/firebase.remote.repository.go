package storage

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/store"
)

type FirebaseRemoteRepository interface {
	UploadFile(ctx context.Context, folder, filename string, fileData []byte, contentType string) (string, error)
	DeleteFile(ctx context.Context, folder, filename string) error
	GetPublicFileURL(folder, filename string) string
}

type firebaseRemoteRepository struct {
	client *store.FirebaseClient
}

func NewFirebaseRemoteRepository(client *store.FirebaseClient) FirebaseRemoteRepository {
	return &firebaseRemoteRepository{
		client: client,
	}
}

func (f *firebaseRemoteRepository) resolveFilePath(folder, filename string) string {
	basePath := os.Getenv("FIREBASE_BASE_PATH")
	profilePhotos := os.Getenv("FIREBASE_PROFILE_PHOTOS")
	return path.Join(basePath, folder, profilePhotos, filename)
}

func (f *firebaseRemoteRepository) UploadFile(ctx context.Context, folder, filename string, fileData []byte, contentType string) (string, error) {
	bucket, err := f.client.StorageClient.Bucket(f.client.BucketName)
	if err != nil {
		return "", fmt.Errorf("error al obtener el bucket")
	}

	fullPath := f.resolveFilePath(folder, filename)
	writer := bucket.Object(fullPath).NewWriter(ctx)
	writer.ContentType = contentType

	// Copiar directamente los datos ya leídos en memoria
	if _, err := writer.Write(fileData); err != nil {
		writer.Close()
		return "", fmt.Errorf("error al copiar archivo a Firebase")
	}

	// Cerrar correctamente el archivo
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error al cerrar la conexión con Firebase")
	}

	// Hace el archivo público
	if err := bucket.Object(fullPath).ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("error al hacer público el archivo")
	}
	return f.GetPublicFileURL(folder, filename), nil
}

// URL publica de la imagen
func (f *firebaseRemoteRepository) GetPublicFileURL(folder, filename string) string {
	baseURLStr := os.Getenv("STORAGE_URL")
	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return ""
	}
	fullPath := path.Join(
		os.Getenv("FIREBASE_BUCKET_NAME"),
		os.Getenv("FIREBASE_BASE_PATH"),
		folder,
		os.Getenv("FIREBASE_PROFILE_PHOTOS"),
		filename,
	)
	baseURL.Path = path.Join(baseURL.Path, fullPath)
	return baseURL.String()
}

// ✅ Eliminar Archivo de Firebase
func (f *firebaseRemoteRepository) DeleteFile(ctx context.Context, folder, filename string) error {
	bucket, err := f.client.StorageClient.Bucket(f.client.BucketName)
	if err != nil {
		return fmt.Errorf("error al obtener el bucket: %w", err)
	}
	fullPath := f.resolveFilePath(folder, filename)

	// 📂 Intentar eliminar el archivo
	object := bucket.Object(fullPath)
	if err := object.Delete(ctx); err != nil {
		return fmt.Errorf("error al eliminar archivo de Firebase: %w", err)
	}

	return nil
}
