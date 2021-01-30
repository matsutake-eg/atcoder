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
        n:usize,
        s:Chars,
    }

    let mut ans = 0;
    for i in 0..n {
        let mut c1: i64 = 0;
        let mut c2: i64 = 0;
        for j in i..n {
            if s[j] == 'A' {
                c1 += 1;
            } else if s[j] == 'T' {
                c1 -= 1;
            } else if s[j] == 'C' {
                c2 += 1;
            } else if s[j] == 'G' {
                c2 -= 1;
            }
            if c1 == 0 && c2 == 0 {
                ans += 1;
            }
        }
    }
    println!("{}", ans);
}
