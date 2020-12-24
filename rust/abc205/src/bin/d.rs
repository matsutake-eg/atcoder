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
        q:usize,
        a:[usize;n],
        k:[usize;q],
    }

    let mut c = vec![0; n];
    for i in 0..n {
        c[i] = a[i] - i - 1;
    }

    for v in k {
        let r = c.lower_bound(&v);
        if r == 0 {
            println!("{}", v);
            continue;
        }
        println!("{}", a[r - 1] + v - c[r - 1]);
    }
}
