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
        n:usize,
        a:[usize;n],
    }

    let mut a = a;
    a.sort();
    for i in 1..=n {
        if i != a[i - 1] {
            println!("No");
            return;
        }
    }
    println!("Yes");
}
