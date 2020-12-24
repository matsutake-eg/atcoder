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
        k:usize,
        a:[usize;n],
    }

    let mut c = vec![0; n + 1];
    for v in a {
        c[v] += 1;
    }

    let mut ans = 0;
    let mut mn = k;
    for v in c {
        mn = min(mn, v);
        ans += mn;
    }
    println!("{}", ans);
}
