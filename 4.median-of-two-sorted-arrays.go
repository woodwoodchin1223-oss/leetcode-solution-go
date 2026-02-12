func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    k := len(nums1) + len(nums2)
    if (len(nums1) + len(nums2)) % 2 == 1 {
        return float64(helper(nums1, nums2, 0, len(nums1) - 1, 0, len(nums2) - 1, k / 2))
    } else {
        r1 := float64(helper(nums1, nums2, 0, len(nums1) - 1, 0, len(nums2) - 1, k / 2))
        r2 := float64(helper(nums1, nums2, 0, len(nums1) - 1, 0, len(nums2) - 1, k / 2 - 1))
        return (r1 + r2) / 2
    }
}

func helper(nums1 []int, nums2 []int, start1 int, end1 int, start2 int, end2 int, k int) int {
    if start1 > end1 {
        return nums2[k - start1]
    }

    if start2 > end2 {
        return nums1[k - start2]
    }

    mid1, mid2 := (start1 + end1) / 2, (start2 + end2) / 2
    v1, v2 := nums1[mid1], nums2[mid2]

    if mid1 + mid2 < k {
        if v1 > v2 {
            return helper(nums1, nums2, start1, end1, mid2 + 1, end2, k)
        } else {
            return helper(nums1, nums2, mid1 + 1, end1, start2, end2, k)
        }
    } else {
        if v1 > v2 {
            return helper(nums1, nums2, start1, mid1 - 1, start2, end2, k)
        } else {
            return helper(nums1, nums2, start1, end1, start2, mid2 - 1, k)
        }
    }
}
