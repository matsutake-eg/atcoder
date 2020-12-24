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
        a:[Chars;h],
    }

    let mut dq = VecDeque::new();
    let mut seen = HashSet::new();
    let mut cnt = vec![vec![0; w]; h];
    for i in 0..h {
        for j in 0..w {
            if a[i][j] == '#' {
                dq.push_back((i, j));
                seen.insert((i, j));
                cnt[i][j] += 1;
            }
        }
    }

    let mut ans = 0;
    while let Some((i, j)) = dq.pop_front() {
        if seen.len() == h * w {
            println!("{}", ans);
            return;
        }
        ans = cnt[i][j];

        let i = i as i64;
        let j = j as i64;
        let h = h as i64;
        let w = w as i64;
        for &(ni, nj) in &[(i, j + 1), (i, j - 1), (i + 1, j), (i - 1, j)] {
            if ni < 0 || ni >= h || nj < 0 || nj >= w {
                continue;
            }
            let ni = ni as usize;
            let nj = nj as usize;
            if seen.contains(&(ni, nj)) {
                continue;
            }
            dq.push_back((ni, nj));
            seen.insert((ni, nj));
            cnt[ni][nj] = ans + 1;
        }
    }
}
