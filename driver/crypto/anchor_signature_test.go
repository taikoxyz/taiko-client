package crypto

import (
	"reflect"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func sig(r, s string, v byte) []byte {
	return append(append(hexutil.MustDecode(r), hexutil.MustDecode(s)...), v)
}

func TestSignAnchor(t *testing.T) {
	type args struct {
		hash []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "k = 1, test case 1",
			args: args{
				hash: hexutil.MustDecode(
					"0x44943399d1507f3ce7525e9be2f987c3db9136dc759cb7f92f742154196868b9",
				),
			},
			want: sig(
				"0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
				"0x782a1e70872ecc1a9f740dd445664543f8b7598c94582720bca9a8c48d6a4766",
				1,
			),
			wantErr: false,
		},
		{
			name: "k = 1, test case 2",
			args: args{
				hash: hexutil.MustDecode(
					"0x663d210fa6dba171546498489de1ba024b89db49e21662f91bf83cdffe788820",
				),
			},
			want: sig(
				"0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
				"0x568130fab1a3a9e63261d4278a7e130588beb51f27de7c20d0258d38a85a27ff",
				1,
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey, err := crypto.Ecrecover(tt.args.hash, tt.want)
			t.Logf("pubKey len %v", len(pubKey))
			require.Nil(t, err)
			verify := crypto.VerifySignature(pubKey, tt.args.hash, tt.want[:64])
			got, err := SignAnchor(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignAnchor() error = %v, wantErr %v, verify %v", err, tt.wantErr, verify)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignAnchor() = %v, want %v, verify %v", got, tt.want, verify)
			} else {
				t.Logf("SignAnchor() = %v, want %v, verify %v", got, tt.want, verify)
			}
		})
	}
}

func TestSignAnchorRS2(t *testing.T) {
	type args struct {
		hash []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "K = 2, test case 1",
			args: args{
				hash: hexutil.MustDecode(
					"0x44943399d1507f3ce7525e9be2f987c3db9136dc759cb7f92f742154196868b9",
				),
			},
			want: sig(
				"0xc6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5",
				"0x38940d69b21d5b088beb706e9ebabe6422307e12863997a44239774467e240d5",
				1,
			),
			wantErr: false,
		},
		{
			name: "K = 2, test case 2",
			args: args{
				hash: hexutil.MustDecode(
					"0x663d210fa6dba171546498489de1ba024b89db49e21662f91bf83cdffe788820",
				),
			},
			want: sig(
				"0xc6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5",
				"0x5840695138a83611aa9dac67beb95aba7323429787a78df993f1c5c7f2c0ef7f",
				0,
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := signWithK(new(secp256k1.ModNScalar).SetInt(2))(tt.args.hash)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignAnchor() = %v, want %v", got, tt.want)
			} else {
				t.Logf("SignAnchor() = %v, want %v", got, tt.want)
			}
		})
	}
}
