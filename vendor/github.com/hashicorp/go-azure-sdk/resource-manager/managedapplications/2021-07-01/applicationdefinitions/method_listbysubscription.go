package applicationdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationDefinition
}

type ListBySubscriptionCompleteResult struct {
	Items []ApplicationDefinition
}

// ListBySubscription ...
func (c ApplicationDefinitionsClient) ListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result ListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Solutions/applicationDefinitions", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ApplicationDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySubscriptionComplete retrieves all the results into a single object
func (c ApplicationDefinitionsClient) ListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (ListBySubscriptionCompleteResult, error) {
	return c.ListBySubscriptionCompleteMatchingPredicate(ctx, id, ApplicationDefinitionOperationPredicate{})
}

// ListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationDefinitionsClient) ListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ApplicationDefinitionOperationPredicate) (result ListBySubscriptionCompleteResult, err error) {
	items := make([]ApplicationDefinition, 0)

	resp, err := c.ListBySubscription(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
