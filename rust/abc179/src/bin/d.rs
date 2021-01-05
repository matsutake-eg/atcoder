#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 998_244_353;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        k:usize,
        lr: [(usize,usize);k],
    }

    let mut dp = vec![0; n + 1];
    dp[1] = 1;
    let mut dp_sum = vec![0; n + 1];
    dp_sum[1] = 1;
    for i in 2..=n {
        for (l, r) in &lr {
            if i < *l {
                continue;
            }
            let ri = i - l;
            let li = if i <= *r { 1 } else { i - r };
            dp[i] += (dp_sum[ri] + MOD) - dp_sum[li - 1];
            dp[i] %= MOD;
        }
        dp_sum[i] = dp[i] + dp_sum[i - 1];
        dp_sum[i] %= MOD;
    }
    println!("{:?}", dp[n]);
}
