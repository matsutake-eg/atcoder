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
    for i in 0..(h - 1) {
        for j in 0..(w - 1) {
            let mut cnt = 0;
            for (x, y) in &[(0, 0), (1, 0), (0, 1), (1, 1)] {
                if s[i + x][j + y] == '#' {
                    cnt += 1;
                }
            }
            if cnt == 1 || cnt == 3 {
                ans += 1;
            }
        }
    }
    println!("{}", ans);
}
