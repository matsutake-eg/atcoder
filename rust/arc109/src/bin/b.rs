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
        n:i128,
    }

    let (mut l, mut r) = (1, n + 1);
    while r - l > 1 {
        let m = (l + r) / 2;
        if (1 + m) * m / 2 > n + 1 {
            r = m;
        } else {
            l = m;
        }
    }
    println!("{}", n + 1 - l);
}
