#![allow(unused_imports)]
use itertools_num::ItertoolsNum as _;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;

#[allow(unused)]
const INF: usize = std::usize::MAX / 4;
#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[fastout]
fn main() {
    input! {
        n:usize,
        a:[usize;n],
    }

    let cumsum = a.iter().cumsum().collect::<Vec<usize>>();

    let mut ans = 0;
    for i in 0..n - 1 {
        ans += a[i] * ((cumsum[n - 1] - cumsum[i]) % MOD);
        ans %= MOD;
    }
    println!("{}", ans);
}
