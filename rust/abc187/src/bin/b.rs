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
        xy: [(i64,i64);n],
    }

    let mut ans = 0;
    for (x1, y1) in &xy {
        for (x2, y2) in &xy {
            if x1 == x2 {
                continue;
            }

            if ((y1 - y2) as f64 / (x1 - x2) as f64).abs() <= 1.0 {
                ans += 1;
            }
        }
    }
    println!("{}", ans / 2);
}
