package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"log"
	"slurp-server/internal/core/ports"
)

type FirestoreApiConfigurationRepository struct {
	ProjectId     string
	ApiCollection string
}

func NewFirestoreRepository(projectId string, collectionPrefix string) FirestoreApiConfigurationRepository {
	return FirestoreApiConfigurationRepository{
		ProjectId:     projectId,
		ApiCollection: collectionPrefix + "api-configuration",
	}
}

func (f FirestoreApiConfigurationRepository) AddApiConfiguration(configuration ports.ApiConfiguration) error {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	defer client.Close()
	api := client.Collection(f.ApiCollection).Doc(configuration.Name)
	_, err = api.Create(ctx, configuration)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

func (f FirestoreApiConfigurationRepository) UpdateApiConfiguration(configuration ports.ApiConfiguration) error {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	api := client.Collection(f.ApiCollection).Doc(configuration.Name)
	_, err = api.Set(ctx, configuration)
	if err != nil {
		return err
	}
	return nil
}

func (f FirestoreApiConfigurationRepository) DeleteApiConfiguration(name string) error {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	api := client.Collection(f.ApiCollection).Doc(name)
	_, err = api.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f FirestoreApiConfigurationRepository) ListApiConfigurations() (*[]ports.ApiConfiguration, error) {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	apis := client.Collection(f.ApiCollection).Documents(ctx)
	var result []ports.ApiConfiguration
	for {
		doc, err := apis.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var apiStruct ports.ApiConfiguration
		if err := doc.DataTo(&apiStruct); err != nil {
			return nil, err
		}
		result = append(result, apiStruct)
	}
	return &result, nil
}

func (f FirestoreApiConfigurationRepository) GetApiConfiguration(name string) (*ports.ApiConfiguration, error) {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	api := client.Collection(f.ApiCollection).Doc(name)
	docSnap, err := api.Get(ctx)
	if err != nil {
		return nil, err
	}
	var apiStruct ports.ApiConfiguration
	if err := docSnap.DataTo(&apiStruct); err != nil {
		return nil, err
	}
	return &apiStruct, nil
}

func (f FirestoreApiConfigurationRepository) getClient(c context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(c, f.ProjectId)
	if err != nil {
		return nil, err
	}
	return client, nil
}
