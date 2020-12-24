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

    let mut p = vec![0; n];
    let mut q = vec![0; n];
    let mut sum = 0;
    let mut mv = 0;
    for i in 0..n {
        sum += a[i];
        p[i] = sum;
        mv = max(mv, sum);
        q[i] = mv;
    }

    let mut ans = 0;
    let mut now = 0;
    for i in 0..n {
        ans = max(ans, now + q[i]);
        now += p[i];
    }
    println!("{}", ans);
}
