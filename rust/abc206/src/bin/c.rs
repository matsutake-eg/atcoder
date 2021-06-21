#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::{fastout, input, marker::*};
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: don't print in 'match arm' because of fastout's bug
#[fastout]
fn main() {
    input! {
        n:usize,
        a:[usize;n],
    }

    let mut hm = HashMap::new();
    for v in a {
        *hm.entry(v).or_insert(0) += 1;
    }

    let mut ans = n * (n - 1) / 2;
    for (_, v) in hm {
        ans -= v * (v - 1) / 2;
    }
    println!("{}", ans);
}
