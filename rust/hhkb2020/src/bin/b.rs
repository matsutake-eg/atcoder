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
        s:[Chars;h],
    }

    let mut ans = 0;
    for i in 0..h {
        for j in 0..w {
            if j + 1 < w {
                if s[i][j] == '.' && s[i][j + 1] == '.' {
                    ans += 1;
                }
            }
            if i + 1 < h {
                if s[i][j] == '.' && s[i + 1][j] == '.' {
                    ans += 1;
                }
            }
        }
    }
    println!("{}", ans);
}
