// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

func Test_IsRootSpan(t *testing.T) {
	exprFunc, err := isRootSpan()
	require.NoError(t, err)

	// root span
	spanRoot := ptrace.NewSpan()
	spanRoot.SetParentSpanID(pcommon.SpanID{
		0, 0, 0, 0, 0, 0, 0, 0,
	})

	rootCtx := ottlspan.BorrowContext(spanRoot, pcommon.NewInstrumentationScope(), pcommon.NewResource(), ptrace.NewScopeSpans(), ptrace.NewResourceSpans())
	ottlspan.ReturnContext(rootCtx)
	value, err := exprFunc(nil, rootCtx)
	require.NoError(t, err)
	require.Equal(t, true, value)

	// non root span
	spanNonRoot := ptrace.NewSpan()
	spanNonRoot.SetParentSpanID(pcommon.SpanID{
		1, 0, 0, 0, 0, 0, 0, 0,
	})

	nonRootCtx := ottlspan.BorrowContext(spanNonRoot, pcommon.NewInstrumentationScope(), pcommon.NewResource(), ptrace.NewScopeSpans(), ptrace.NewResourceSpans())
	ottlspan.ReturnContext(nonRootCtx)
	value, err = exprFunc(nil, nonRootCtx)
	require.NoError(t, err)
	require.Equal(t, false, value)
}
