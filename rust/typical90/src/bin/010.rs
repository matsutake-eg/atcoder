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
        cp: [(usize,usize);n],
        q:usize,
        lr: [(usize,usize);q],
    }

    let mut sum1 = vec![0; n + 1];
    let mut sum2 = vec![0; n + 1];
    for (i, (c, p)) in cp.into_iter().enumerate() {
        sum1[i + 1] = sum1[i];
        sum2[i + 1] = sum2[i];
        if c == 1 {
            sum1[i + 1] += p;
        } else {
            sum2[i + 1] += p;
        }
    }

    for (l, r) in lr {
        println!("{} {}", sum1[r] - sum1[l - 1], sum2[r] - sum2[l - 1]);
    }
}
