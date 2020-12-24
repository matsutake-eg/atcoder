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
        h:usize,
        w:usize,
        n:usize,
        m:usize,
        ab: [(Usize1,Usize1);n],
        cd: [(Usize1,Usize1);m],
    }

    let mut light = vec![vec![false; w]; h];
    for (a, b) in ab {
        light[a][b] = true;
    }
    let mut block = vec![vec![false; w]; h];
    for (c, d) in cd {
        block[c][d] = true;
    }

    let mut akari = vec![vec![false; w]; h];
    for i in 0..h {
        let mut is_lighten = false;
        for j in 0..w {
            if light[i][j] {
                is_lighten = true;
            }
            if block[i][j] {
                is_lighten = false;
            }
            if is_lighten {
                akari[i][j] = true;
            }
        }

        let mut is_lighten = false;
        for j in (0..w).rev() {
            if light[i][j] {
                is_lighten = true;
            }
            if block[i][j] {
                is_lighten = false;
            }
            if is_lighten {
                akari[i][j] = true;
            }
        }
    }
    for j in 0..w {
        let mut is_lighten = false;
        for i in 0..h {
            if light[i][j] {
                is_lighten = true;
            }
            if block[i][j] {
                is_lighten = false;
            }
            if is_lighten {
                akari[i][j] = true;
            }
        }

        let mut is_lighten = false;
        for i in (0..h).rev() {
            if light[i][j] {
                is_lighten = true;
            }
            if block[i][j] {
                is_lighten = false;
            }
            if is_lighten {
                akari[i][j] = true;
            }
        }
    }
    let mut ans = 0;
    for row in akari {
        for v in row {
            if v {
                ans += 1;
            }
        }
    }
    println!("{}", ans);
}
