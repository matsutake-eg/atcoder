#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        a:i64,
        b:i64,
        c:i64,
    }

    if c % 2 == 1 {
        solve(a, b);
    } else {
        solve(a * a, b * b);
    }
}

fn solve(x: i64, y: i64) {
    if x == y {
        println!("=");
    } else if x < y {
        println!("<");
    } else {
        println!(">");
    }
}
