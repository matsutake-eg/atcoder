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
        n:usize,
        a:[usize;n],
    }

    let mut ans = 0;
    for i in 0..n {
        let mut mn = std::usize::MAX;
        for j in i..n {
            mn = min(mn, a[j]);
            ans = max(ans, mn * (j - i + 1));
        }
    }
    println!("{}", ans);
}
