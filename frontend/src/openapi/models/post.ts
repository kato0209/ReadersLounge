/* tslint:disable */
/* eslint-disable */
/**
 * ReadersLounge API
 * ReadersLounge API
 *
 * The version of the OpenAPI document: 1.0.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

// May contain unused imports in some cases
// @ts-ignore
import { Book } from './book';
// May contain unused imports in some cases
// @ts-ignore
import { PostLike } from './post-like';
// May contain unused imports in some cases
// @ts-ignore
import { User } from './user';

/**
 *
 * @export
 * @interface Post
 */
export interface Post {
  /**
   *
   * @type {number}
   * @memberof Post
   */
  post_id: number;
  /**
   *
   * @type {User}
   * @memberof Post
   */
  user: User;
  /**
   *
   * @type {string}
   * @memberof Post
   */
  content: string;
  /**
   *
   * @type {number}
   * @memberof Post
   */
  rating: number;
  /**
   *
   * @type {string}
   * @memberof Post
   */
  image?: string;
  /**
   *
   * @type {string}
   * @memberof Post
   */
  created_at: string;
  /**
   *
   * @type {Book}
   * @memberof Post
   */
  book: Book;
  /**
   *
   * @type {Array<PostLike>}
   * @memberof Post
   */
  likes?: Array<PostLike>;
}
