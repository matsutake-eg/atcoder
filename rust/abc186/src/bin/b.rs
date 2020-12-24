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
        h:usize,
        w:usize,
        a:[[usize;w];h],
    }

    let mut mn = std::usize::MAX;
    for row in &a {
        for v in row {
            mn = min(mn, *v);
        }
    }

    let mut ans = 0;
    for row in a {
        for v in row {
            ans += v - mn;
        }
    }
    println!("{}", ans);
}
