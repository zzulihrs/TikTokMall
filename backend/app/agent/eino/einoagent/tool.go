/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package einoagent

import (
	"context"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/cart"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/einotool"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/gitclone"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/open"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/orderlookup"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/payment"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/productlist"
	"github.com/tiktokmall/backend/app/agent/pkg/tool/task"
)

func GetTools(ctx context.Context) ([]tool.BaseTool, error) {
	einoAssistantTool, err := NewEinoAssistantTool(ctx)
	if err != nil {
		return nil, err
	}

	toolTask, err := NewTaskTool(ctx)
	if err != nil {
		return nil, err
	}

	toolOpen, err := NewOpenFileTool(ctx)
	if err != nil {
		return nil, err
	}

	toolGitClone, err := NewGitCloneFile(ctx)
	if err != nil {
		return nil, err
	}

	// toolDDGSearch, err := NewDDGSearch(ctx, nil)
	// if err != nil {
	// 	return nil, err
	// }

	toolOrderLookUp, err := NewOrderLookUpTool(ctx)
	if err != nil {
		return nil, err
	}

	toolProductList, err := NewProdcutListTool(ctx)
	if err != nil {
		return nil, err
	}

	toolCart, err := NewCartTool(ctx)
	if err != nil {
		return nil, err
	}

	toolPayment, err := NewPaymentTool(ctx)
	if err != nil {
		return nil, err
	}
	return []tool.BaseTool{
		einoAssistantTool,
		toolTask,
		toolOpen,
		toolGitClone,
		// toolDDGSearch,
		toolOrderLookUp,
		toolProductList,
		toolCart,
		toolPayment,
	}, nil
}

func defaultDDGSearchConfig(ctx context.Context) (*duckduckgo.Config, error) {
	config := &duckduckgo.Config{}
	return config, nil
}

func NewDDGSearch(ctx context.Context, config *duckduckgo.Config) (tn tool.BaseTool, err error) {
	if config == nil {
		config, err = defaultDDGSearchConfig(ctx)
		if err != nil {
			return nil, err
		}
	}
	tn, err = duckduckgo.NewTool(ctx, config)
	if err != nil {
		return nil, err
	}
	return tn, nil
}

func NewOpenFileTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return open.NewOpenFileTool(ctx, nil)
}

func NewGitCloneFile(ctx context.Context) (tn tool.BaseTool, err error) {
	return gitclone.NewGitCloneFile(ctx, nil)
}

func NewEinoAssistantTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return einotool.NewEinoAssistantTool(ctx, nil)
}

func NewTaskTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return task.NewTaskTool(ctx, nil)
}

func NewOrderLookUpTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return orderlookup.NewOrderLookup(ctx, nil)
}

func NewProdcutListTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return productlist.NewProduct(ctx, nil)
}

func NewCartTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return cart.NewCartTool(ctx, nil)
}

func NewPaymentTool(ctx context.Context) (tn tool.BaseTool, err error) {
	return payment.NewPaymentTool(ctx, nil)
}
