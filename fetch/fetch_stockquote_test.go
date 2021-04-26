package fetch

import (
	"io"
	"reflect"
	"testing"

	"github.com/dude333/rapina"
)

type MockStockFetch struct {
	// store rapina.StockStore
}

func (m MockStockFetch) CsvToDB(stream io.ReadCloser, code string) error {
	return nil
}

func (m MockStockFetch) Quote(code, date string) (rapina.Quotation, error) {
	// fmt.Printf("calling mock Quote(%s, %s)\n", code, date)
	return rapina.Quotation{
		Code: code,
		Date: date,
		Val:  123.45,
	}, nil
}

func TestStockFetch_Quote(t *testing.T) {
	type fields struct {
		apiKey string
		store  rapina.StockStore
	}
	type args struct {
		code string
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    rapina.Quotation
		wantErr bool
	}{
		{
			name: "should return a quote",
			fields: fields{
				apiKey: "test",
				store:  MockStockFetch{},
			},
			args: args{
				code: "TEST11",
				date: "2021-04-26",
			},
			want: rapina.Quotation{
				Code: "TEST11",
				Date: "2021-04-26",
				Val:  123.45,
			},
			wantErr: false,
		},
		{
			name: "should return a date error",
			fields: fields{
				apiKey: "test",
				store:  MockStockFetch{},
			},
			args: args{
				code: "TEST11",
				date: "2021-04-32",
			},
			want: rapina.Quotation{
				Code: "TEST11",
				Date: "2021-04-26",
				Val:  123.45,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StockFetch{
				apiKey: tt.fields.apiKey,
				store:  tt.fields.store,
			}
			got, err := s.Quote(tt.args.code, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("StockFetch.Quote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StockFetch.Quote() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
