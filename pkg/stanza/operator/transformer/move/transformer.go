// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package move // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/move"

import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that moves a field's value to a new field
type Transformer struct {
	helper.TransformerOperator
	From entry.Field
	To   entry.Field
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	return t.ProcessBatchWith(ctx, entries, t.process)
}

// Transform will apply the move operation to an entry
func (t *Transformer) process(_ context.Context, e *entry.Entry) error {
	val, exist := t.From.Delete(e)
	if !exist {
		return fmt.Errorf("move: field does not exist: %s", t.From.String())
	}
	return t.To.Set(e, val)
}
