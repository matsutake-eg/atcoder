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
        h:usize,
        w:usize,
        a:[[usize;w];h],
    }

    let mut hsum = vec![0; h];
    let mut wsum = vec![0; w];
    for i in 0..h {
        for j in 0..w {
            hsum[i] += a[i][j];
            wsum[j] += a[i][j];
        }
    }

    for i in 0..h {
        for j in 0..w {
            print!("{}", hsum[i] + wsum[j] - a[i][j]);
            if j == w - 1 {
                println!("");
            } else {
                print!(" ");
            }
        }
    }
}
