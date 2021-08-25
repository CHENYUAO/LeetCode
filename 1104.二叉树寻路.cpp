/*
 * @lc app=leetcode.cn id=1104 lang=cpp
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
class Solution {
public:
int findRow(int label) {
	for (int i = 0;; i++) {
		if (pow(2, i) <= label && label < pow(2, i + 1)) {
			return i;
		}
	}
}

vector<int> pathInZigZagTree(int label) {
    int row = findRow(label);
    vector<int> res{ label };
    for (int i = row; i > 0; i--) {
	    int sum = pow(2, i - 1) + pow(2, i) - 1;
	    label = sum - label / 2;
	    res.push_back(label);
	    /*cout << "sum:" << sum << "   label:" << label << endl;*/
    }
    reverse(res.begin(), res.end());
    return res;
}
};
// @lc code=end

