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

    for (x1, y1) in &xy {
        for (x2, y2) in &xy {
            if x1 == x2 && y1 == y2 {
                continue;
            }
            for (x3, y3) in &xy {
                if (x3 == x1 && y3 == y1) || (x3 == x2 && y3 == y2) {
                    continue;
                }
                if (y2 - y1) * (x3 - x2) == (y3 - y2) * (x2 - x1) {
                    println!("Yes");
                    return;
                }
            }
        }
    }
    println!("No");
}
