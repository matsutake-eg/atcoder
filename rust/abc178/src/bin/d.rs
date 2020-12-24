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
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        s:usize,
    }

    let mut dp = vec![0; s + 1];
    dp[0] = 1;
    for i in 3..=s {
        for j in 0..=(i - 3) {
            dp[i] += dp[j];
            dp[i] %= MOD;
        }
    }
    println!("{:?}", dp[s]);
}
