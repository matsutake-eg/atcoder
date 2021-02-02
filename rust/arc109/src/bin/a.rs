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
        a:i64,
        b:i64,
        x:i64,
        y:i64,
    }

    let d = (a - b).abs();
    if a <= b {
        println!("{}", x + min(x * 2, y) * d);
    } else {
        println!("{}", x + min(x * 2, y) * (d - 1));
    }
}
