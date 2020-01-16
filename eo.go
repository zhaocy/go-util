package go_util

import "math"

// 算法计算
// RA RB 段位分
func Elo_rating_system(RA, RB, ScoreA, ScoreB int) (A, B float64) {
	var K = float64(32)
	var EA = 1 / (1 + math.Pow(float64(10), float64(RB-RA)/400))
	var EB = 1 / (1 + math.Pow(float64(10), float64(RA-RB)/400))
	var SA = Compare(ScoreA, ScoreB)
	var SB = Compare(ScoreB, ScoreA)

	var R_A_L = K * (SA - EA)
	var R_B_L = K * (SB - EB)

	R_A_L = float64(RA) + Accuracy(R_A_L);
	R_B_L = float64(RB) + Accuracy(R_B_L);
	var R_A_K = Accuracy(K * (SA - EA));
	var R_B_K = Accuracy(K * (SB - EB));
	return float64(RA) + R_A_K, float64(RB) + R_B_K;

}

// 计算精度（当计算的绝对值小于1时，取1）
func Accuracy(value float64) float64 {
	var r = math.Abs(value);
	if (r > 0 && r < 1) {
		r = 1
	}
	if (value >= 0) {
		return math.Round(r);
	} else {
		return -math.Round(r);
	}
}

// 游戏积分结果比较
func Compare(ScoreA, ScoreB int) float64 {
	if (ScoreA > ScoreB) {
		return 1;
	} else if (ScoreA == ScoreB) {
		return 0.5;
	} else {
		return 0
	}
}
