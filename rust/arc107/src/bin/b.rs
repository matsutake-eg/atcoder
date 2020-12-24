#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:i64,
        k:i64,
    }

    let mut ans = 0;
    for i in 2..=n * 2 {
        if i + k >= 2 && i + k <= n * 2 {
            ans += min(i - 1, n * 2 + 1 - i) * min(i + k - 1, n * 2 + 1 - i - k);
        }
    }
    println!("{}", ans);
}
