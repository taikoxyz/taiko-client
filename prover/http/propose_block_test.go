package http

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyberhorsey/webutils/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	taikotestutils "github.com/taikoxyz/taiko-client/testutils"
)

func Test_ProposeBlock(t *testing.T) {
	srv := newTestServer("")

	tests := []struct {
		name                  string
		req                   *encoding.ProposeBlockData
		wantStatus            int
		wantBodyRegexpMatches []string
	}{
		{
			"success",
			&encoding.ProposeBlockData{
				Fee:    big.NewInt(1000),
				Expiry: uint64(time.Now().Unix()),
				Input: encoding.TaikoL1BlockMetadataInput{
					Beneficiary:     common.BytesToAddress(taikotestutils.RandomHash().Bytes()),
					TxListHash:      taikotestutils.RandomHash(),
					TxListByteStart: common.Big0,
					TxListByteEnd:   common.Big0,
					CacheTxListInfo: false,
				},
			},
			http.StatusOK,
			[]string{`"signedPayload"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := testutils.NewUnauthenticatedRequest(
				echo.POST,
				fmt.Sprintf("/proposeBlock"),
				tt.req,
			)

			rec := httptest.NewRecorder()

			srv.ServeHTTP(rec, req)

			testutils.AssertStatusAndBody(t, rec, tt.wantStatus, tt.wantBodyRegexpMatches)
		})
	}
}
