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
        h:usize,
        w:usize,
        a:[[usize;w];h],
    }

    let mut row = vec![0; h];
    let mut col = vec![0; w];
    for i in 0..h {
        for j in 0..w {
            row[i] += a[i][j];
            col[j] += a[i][j];
        }
    }

    let mut ans = vec![vec![0; w]; h];
    for i in 0..h {
        for j in 0..w {
            ans[i][j] = row[i] + col[j] - a[i][j];
        }
    }

    for i in 0..h {
        for j in 0..w {
            print!("{}", ans[i][j]);
            if j < w - 1 {
                print!(" ");
            }
        }
        println!("");
    }
}
