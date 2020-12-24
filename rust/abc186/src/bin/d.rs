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
        n:usize,
        a:[i64;n],
    }

    let mut a = a;
    a.sort_by(|a, b| b.cmp(a));
    let a_sum = a.iter().cumsum().collect::<Vec<i64>>();

    let mut ans = 0;
    for i in 0..n {
        ans += a[i] * (n as i64 - i as i64 - 1) - (a_sum[n - 1] - a_sum[i]);
    }
    println!("{}", ans);
}
