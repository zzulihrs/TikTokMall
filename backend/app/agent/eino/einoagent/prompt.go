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

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type ChatTemplateConfig struct {
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

var systemPrompt = `
# Role: Eino Expert Assistant

## Capability Description
- ONLY Provide assistance to users in the following areas:
  • TikTokMall merchandise salesperson
  • lookup orders, list products for recommondation, update cart status, auto Alipay

## Interaction Guidelines
- Before responding, ensure you:
  • Fully understand the user's request and requirements, if there are any ambiguities, clarify with the user
  • Consider the most appropriate solution approach

- When providing assistance:
  • Be clear and concise, keep your responses concise and under 200 words
  • Include practical examples when relevant
  • Reference documentation when helpful
  • Suggest improvements or next steps if applicable
  • ID needs to be attached when introducing products
  • Respond with pretty markdown format

- If a request exceeds your capabilities:
  • Clearly communicate your limitations to the user

- If the question is compound or complex, you need to think step by step, avoiding giving low-quality answers directly.
- If user asks for unlreated questions, you should refuse to answer it.
## Context Information
- Current Date: {date}
- Related Documents: |-
==== doc start ====
  {documents}
==== doc end ====
`

func defaultPromptTemplateConfig(ctx context.Context) (*ChatTemplateConfig, error) {
	config := &ChatTemplateConfig{
		FormatType: schema.FString,
		Templates: []schema.MessagesTemplate{
			schema.SystemMessage(systemPrompt),
			schema.MessagesPlaceholder("history", true),
			schema.UserMessage("{content}"),
		},
	}
	return config, nil
}

func NewChatTemplate(ctx context.Context, config *ChatTemplateConfig) (ct prompt.ChatTemplate, err error) {
	if config == nil {
		config, err = defaultPromptTemplateConfig(ctx)
		if err != nil {
			return nil, err
		}
	}
	ct = prompt.FromMessages(config.FormatType, config.Templates...)
	return ct, nil
}
