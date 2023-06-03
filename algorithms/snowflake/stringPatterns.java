This is a dynamic programming problem. Let's imagine the tree of unique words to be generated.
wordLen = 3 and maxVowels = 1
There are two choices:

Start with consonant 'c' or vowel 'v'(if maxVowels > 0)
If we start with 'c' then the next option is either 'c' or 'v'
If we start with 'v' then the next option is either 'c'. The option 'v' is not possible we have max consecutive vowels equals to 1.
So we build the tree with possible ways to generate unique words. Let's mark the tree nodes with state. The top starts with (3,1). That means the node have length is 3 and max consecutive vowels is 1.
The child nodes of top are (2,1) and (2,0) and etc. When we reach the node (2,0) it's 'v' there is only one child node. It's (1,1) 'c'. We should start with maxVowels since there no more vowels before.
The idea is to calculate each subtree of the tree we built quickly. We need to memorize intermediate results. And reuse them in further calculations.

There are two solutions: dfs with memorization(top-down) and bfs(bottom-up):

int[][][] memo;
static int mod = (int)Math.pow(10, 9) + 7;
// TC worldLen * vowels, SC worldLen * vowels
public int solveWithMemo(int wordLen, int vowels) {
memo = new int[wordLen + 1][vowels+1][2];
return (solveRecursively(wordLen - 1, vowels, 1) + solveRecursively(wordLen - 1, vowels - 1, 0)) % mod;
}

/**
 * @param charType 1 consonant or 0 vowel
 */
public int solveRecursively(int wordLen, int vowels, int charType) {
    if (vowels < 0 || wordLen < 0)
        return 0;
    int v = charType == 1 ? 21 : 5;
    if (wordLen == 0) {
        return v;
    }
    if (memo[wordLen][vowels][charType] != 0)
        return memo[wordLen][vowels][charType];
    int res;
    if (vowels == 0) {
        long res1 = solveRecursively(wordLen - 1, maxVowels, 1);
        res = (int)((res1 * v) % mod);
    } else {
        long res1 = solveRecursively(wordLen - 1, vowels, 1);
        long res2 = solveRecursively(wordLen - 1, vowels - 1, 0);
        res = (int)(((res1 + res2) % mod) * v % mod);
    }
    memo[wordLen][vowels][charType] = res;
    return res;
}