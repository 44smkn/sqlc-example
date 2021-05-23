package config_test

import (
	"testing"

	"github.com/44smkn/sqlc-sample/pkg/config"
	"github.com/google/go-cmp/cmp"
)

func TestGetChairCondtion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want config.ChairSearchCondition
	}{
		{
			name: "simple",
			want: config.ChairSearchCondition{
				Width: config.RangeCondition{
					Prefix: "",
					Suffix: "cm",
					Ranges: []*config.Range{
						{Min: -1, Max: 80},
						{ID: 1, Min: 80, Max: 110},
						{ID: 2, Min: 110, Max: 150},
						{ID: 3, Min: 150, Max: -1},
					},
				},
				Height: config.RangeCondition{
					Prefix: "",
					Suffix: "cm",
					Ranges: []*config.Range{
						{Min: -1, Max: 80},
						{ID: 1, Min: 80, Max: 110},
						{ID: 2, Min: 110, Max: 150},
						{ID: 3, Min: 150, Max: -1},
					},
				},
				Depth: config.RangeCondition{
					Prefix: "",
					Suffix: "cm",
					Ranges: []*config.Range{
						{Min: -1, Max: 80},
						{ID: 1, Min: 80, Max: 110},
						{ID: 2, Min: 110, Max: 150},
						{ID: 3, Min: 150, Max: -1},
					},
				},
				Price: config.RangeCondition{
					Prefix: "",
					Suffix: "円",
					Ranges: []*config.Range{
						{Min: -1, Max: 3000},
						{ID: 1, Min: 3000, Max: 6000},
						{ID: 2, Min: 6000, Max: 9000},
						{ID: 3, Min: 9000, Max: 12000},
						{ID: 4, Min: 12000, Max: 15000},
						{ID: 5, Min: 15000, Max: -1},
					},
				},
				Color: config.ListCondition{
					List: []string{
						"黒", "白", "赤", "青", "緑", "黄", "紫", "ピンク", "オレンジ",
						"水色", "ネイビー", "ベージュ",
					},
				},
				Feature: config.ListCondition{
					List: []string{
						"ヘッドレスト付き", "肘掛け付き", "キャスター付き",
						"アーム高さ調節可能", "リクライニング可能",
						"高さ調節可能", "通気性抜群", "メタルフレーム", "低反発",
						"木製", "背もたれつき", "回転可能", "レザー製", "昇降式",
						"デザイナーズ", "金属製", "プラスチック製", "法事用", "和風", "中華風", "西洋風", "イタリア製",
						"国産", "背もたれなし", "ラテン風", "布貼地", "スチール製", "メッシュ貼地", "オフィス用", "料理店用",
						"自宅用", "キャンプ用", "クッション性抜群", "モーター付き", "ベッド一体型", "ディスプレイ配置可能", "ミニ机付き",
						"スピーカー付属", "中国製", "アンティーク", "折りたたみ可能", "重さ500g以内", "24回払い無金利", "現代的デザイン",
						"近代的なデザイン", "ルネサンス的なデザイン", "アームなし", "オーダーメイド可能", "ポリカーボネート製", "フットレスト付き",
					},
				},
				Kind: config.ListCondition{
					List: []string{
						"ゲーミングチェア", "座椅子", "エルゴノミクス", "ハンモック",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.GetChairSearchCondition()
			if err != nil {
				t.Errorf("getChairSearchCondtion() is failed: %v", err)
			}
			if diff := cmp.Diff(tt.want, *got); diff != "" {
				t.Errorf("getChairSearchCondtion() mismatch (-want +got):\n%s", diff)
			}
		})
	}

}
