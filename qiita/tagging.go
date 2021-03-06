package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Represents an association between an item and a tag.
type Tagging struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}

type Taggings []Tagging

/*
	Add a tag to an item (only available on Qiita:Team)

	POST /api/v2/items/:item_id/taggings
*/
func (c *Client) AddItemTagging(ctx context.Context, itemId string, tagging Tagging) error {
	b, _ := json.Marshal(tagging)
	p := fmt.Sprintf("/api/v2/items/%s/taggings", itemId)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Remove a tag from an item (only available on Qiita:Team)

	DELETE /api/v2/items/:item_id/taggings/:tagging_id
*/
func (c *Client) DeleteItemTagging(ctx context.Context, itemId, taggingId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/taggings/%s", itemId, taggingId)
	res, err := c.delete(ctx, p)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
