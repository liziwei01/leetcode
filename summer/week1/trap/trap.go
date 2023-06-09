/*
 * @Author: liziwei01
 * @Date: 2023-06-06 14:31:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-06-06 18:25:23
 * @Description: file content
 */
package trap

func trap(height []int) int {
	// 剪枝
	if len(height) <= 2 {
		return 0
	}
	// 两端添加0以保证可以获得两端的峰值
	height = append([]int{0}, height...)
	height = append(height, 0)
	// 找到所有的峰值
	peaks := findPeaks(height)
	if len(peaks) < 2 {
		return 0
	}
	// 去除两个峰值之间所有比两段小的峰值
	highPeaks := removeSmallPeaks(height, peaks)
	// 计算存水量
	vol := getVolumn(height, highPeaks)
	return vol
}

func findPeaks(height []int) []int {
	peaks := make([]int, 0)
	raise := false
	// 上一位状态的值
	idx := 0

	// 三状态，平，升，降。检查到升时记录下来。检查到降时且上一个状态是升则视为峰值
	for i := 0; i != len(height); i++ {
		if idx == height[i] {
			// do nothing
		} else if idx < height[i] {
			raise = true
		} else {
			if raise == true {
				peaks = append(peaks, i-1)
			}
			raise = false
		}
		// 更新上一位
		idx = height[i]
	}
	return peaks
}

func removeSmallPeaks(height, peaks []int) []int {
	lIdx := 0
	rIdx := len(peaks) - 1

	for lIdx < rIdx {
		lVal := height[peaks[lIdx]]
		rVal := height[peaks[rIdx]]
		minVal := minInt(lVal, rVal)
		for i := lIdx + 1; i < rIdx; i++ {
			if height[peaks[i]] <= minVal {
				peaks = append(peaks[:i], peaks[i+1:]...)
				rIdx--
				i--
			}
		}
		if lVal < rVal {
			lIdx++
		} else {
			rIdx--
		}
	}

	return peaks
}

func getVolumn(height, peaks []int) int {
	vol := 0

	for i := 0; i != len(peaks) - 1; i++ {
		peak1 := peaks[i]
		peak2 := peaks[i+1]
		minPeakHeight := height[peak2]
		if height[peak1] < height[peak2] {
			minPeakHeight = height[peak1]
		}
		for j := peak1+1; j < peak2; j++ {
			if height[j] < minPeakHeight {
				vol = vol + (minPeakHeight - height[j])
			}
		}
	}

	return vol
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
