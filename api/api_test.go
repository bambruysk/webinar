package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"webinar/api/mocks"
)

func Test_api_handle(t *testing.T) {
	type fields struct {
		r       *mux.Router
		service *mocks.Multier
	}
	type args struct {
		value string
	}
	type want struct {
		statusCode      int
		value           string
		wantCallService bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Test 1",
			fields: fields{
				r: mux.NewRouter(),
				/// обратите внимание вставляем вместе сервиса моковую структуру
				service: &mocks.Multier{},
			},
			args: args{
				value: "2",
			},
			want: want{
				statusCode:      http.StatusOK,
				value:           "4",
				wantCallService: true,
			},
		},
		{
			name: "Test Not number",
			fields: fields{
				r: mux.NewRouter(),
				/// обратите внимание вставляем вместе сервиса моковую структуру
				service: &mocks.Multier{},
			},
			args: args{
				value: "bumbadabum",
			},
			want: want{
				statusCode:      http.StatusBadRequest,
				value:           "",
				wantCallService: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//  если в процессе теста выщываается сервси мокаем
			if tt.want.wantCallService {
				val, err := strconv.Atoi(tt.args.value)
				assert.NoError(t, err)
				res, err := strconv.Atoi(tt.want.value)
				assert.NoError(t, err)
				tt.fields.service.On("Mul2", val).Return(res)
			}

			a := &api{
				r:       tt.fields.r,
				service: tt.fields.service,
			}

			url := fmt.Sprintf("http://%s/%s", httptest.DefaultRemoteAddr, tt.args.value)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			// только для гориллы - добавляем параметр
			params := map[string]string{
				"num": tt.args.value, // важно чтобы имя совпадало с тем, что вытаскивается в обработчике
			}

			r := mux.SetURLVars(request, params)

			h := http.HandlerFunc(a.handle)
			h.ServeHTTP(w, r)
			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.want.statusCode {
				t.Errorf("Status code mismatched got %v, want %v", res.StatusCode, tt.want.statusCode)
			}

			{
				result, err := ioutil.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.value, string(result))
			}
			// проверяем что метод у мокового сервиса вызывался
			if tt.want.wantCallService {
				val, err := strconv.Atoi(tt.args.value)
				assert.NoError(t, err)
				tt.fields.service.AssertCalled(t, "Mul2", val)
				tt.fields.service.AssertNumberOfCalls(t, "Mul2", 1)
			}

		})
	}
}
