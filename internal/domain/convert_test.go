package domain

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
)

func TestModel_Convert(t *testing.T) {
	type currencyPairRepoMockFunc func(mc *minimock.Controller) CurrencyPairRepo

	var (
		mc = minimock.NewController(t)
	)

	type args struct {
		ctx          context.Context
		currencyFrom string
		currencyTo   string
		value        float64
	}
	tests := []struct {
		name             string
		args             args
		want             *ConvertResult
		err              error
		currencyPairRepo currencyPairRepoMockFunc
	}{
		{
			name: "success convert",
			args: args{
				currencyFrom: "USD",
				currencyTo:   "RUB",
				value:        3,
			},
			want: &ConvertResult{
				Well:  70,
				Value: 210,
			},
			currencyPairRepo: func(mc *minimock.Controller) CurrencyPairRepo {
				mock := NewCurrencyPairRepoMock(mc)
				mock.GetMock.Set(func(ctx context.Context, currencyFrom string, currencyTo string) (*CurrencyPair, error) {
					return &CurrencyPair{
						CurrencyFrom: "USD",
						CurrencyTo:   "RUB",
						Well:         70,
						UpdatedAt:    time.Now(),
					}, nil
				})

				return mock
			},
		},
		{
			name: "error convert: not found currency pair",
			args: args{
				currencyFrom: "USD",
				currencyTo:   "RUB",
				value:        3,
			},
			err: CurrencyPairNotFound,
			currencyPairRepo: func(mc *minimock.Controller) CurrencyPairRepo {
				mock := NewCurrencyPairRepoMock(mc)
				mock.GetMock.Set(func(ctx context.Context, currencyFrom string, currencyTo string) (*CurrencyPair, error) {
					return nil, CurrencyPairNotFound
				})

				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				currencyPairRepo: tt.currencyPairRepo(mc),
			}
			got, err := m.Convert(tt.args.ctx, tt.args.currencyFrom, tt.args.currencyTo, tt.args.value)
			if (err != nil) && err == tt.err {
				t.Errorf("Model.Convert() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
