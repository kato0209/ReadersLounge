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

/**
 *
 * @export
 * @interface BookGenreNode
 */
export interface BookGenreNode {
  /**
   *
   * @type {number}
   * @memberof BookGenreNode
   */
  id: number;
  /**
   *
   * @type {string}
   * @memberof BookGenreNode
   */
  books_genre_id: string;
  /**
   *
   * @type {string}
   * @memberof BookGenreNode
   */
  books_genre_name: string;
  /**
   *
   * @type {number}
   * @memberof BookGenreNode
   */
  genre_level: number;
  /**
   *
   * @type {string}
   * @memberof BookGenreNode
   */
  parent_genre_id: string;
  /**
   *
   * @type {Array<BookGenreNode>}
   * @memberof BookGenreNode
   */
  children: Array<BookGenreNode>;
}
