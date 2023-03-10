package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"slurp-server/internal/core/ports"
)

type FirestoreApiHistoryRepository struct {
	ProjectId         string
	HistoryCollection string
}

func NewFirestoreHistoryRepository(projectId string, collectionPrefix string) FirestoreApiHistoryRepository {
	return FirestoreApiHistoryRepository{
		ProjectId:         projectId,
		HistoryCollection: collectionPrefix + "history",
	}
}

func (f FirestoreApiHistoryRepository) AddApiHistory(history ports.ApiHistory) error {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	api := client.Collection(f.HistoryCollection).Doc(history.Name + fmt.Sprintf("%d", history.Date.UnixMilli()))
	_, err = api.Create(ctx, history)
	if err != nil {
		return err
	}
	return nil
}

func (f FirestoreApiHistoryRepository) DeleteApiHistory(id string) error {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	api := client.Collection(f.HistoryCollection).Doc(id)
	_, err = api.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f FirestoreApiHistoryRepository) ListApiHistories() (*[]ports.ApiHistory, error) {
	ctx := context.Background()
	client, err := f.getClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	apis := client.Collection(f.HistoryCollection).Documents(ctx)
	var result []ports.ApiHistory
	for {
		doc, err := apis.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var apiStruct ports.ApiHistory
		if err := doc.DataTo(&apiStruct); err != nil {
			return nil, err
		}
		result = append(result, apiStruct)
	}
	return &result, nil
}

func (f FirestoreApiHistoryRepository) getClient(c context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(c, f.ProjectId)
	if err != nil {
		return nil, err
	}
	return client, nil
}
