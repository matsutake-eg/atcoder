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
        n_max:i64,
        m:usize,
        t:i64,
        ab: [(i64,i64);m],
    }

    let mut n = n_max;
    let mut p = 0;
    for (a, b) in ab {
        n -= a - p;
        if n <= 0 {
            println!("No");
            return;
        }
        n += b - a;
        n = min(n, n_max);
        p = b;
    }

    n -= t - p;
    println!("{}", if n <= 0 { "No" } else { "Yes" });
}
