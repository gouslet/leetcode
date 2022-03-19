/*
 * File: \powerOfTwo\powerOfTwo.go                                             *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2022/03/19 , 14:19:13                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 16:16:57                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
