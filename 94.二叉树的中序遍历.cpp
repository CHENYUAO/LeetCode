/*
 * @lc app=leetcode.cn id=94 lang=cpp
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> stk;
        TreeNode* curr = root;
        // 左节点push进栈
        while (curr!=NULL || !stk.empty()){
            while(curr != NULL){
                stk.push(curr);
                curr = curr->left;
            }
            // 出栈
            curr = stk.top();
            stk.pop();
            res.push_back(curr->val);
            // 右子树进栈
            curr = curr->right;
        }
        return res; 
    }
};
// @lc code=end

