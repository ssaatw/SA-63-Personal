/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntJobtitleEdges,
    EntJobtitleEdgesFromJSON,
    EntJobtitleEdgesFromJSONTyped,
    EntJobtitleEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntJobtitle
 */
export interface EntJobtitle {
    /**
     * Jobtitlename holds the value of the "Jobtitlename" field.
     * @type {string}
     * @memberof EntJobtitle
     */
    jobtitlename?: string;
    /**
     * 
     * @type {EntJobtitleEdges}
     * @memberof EntJobtitle
     */
    edges?: EntJobtitleEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntJobtitle
     */
    id?: number;
}

export function EntJobtitleFromJSON(json: any): EntJobtitle {
    return EntJobtitleFromJSONTyped(json, false);
}

export function EntJobtitleFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntJobtitle {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'jobtitlename': !exists(json, 'Jobtitlename') ? undefined : json['Jobtitlename'],
        'edges': !exists(json, 'edges') ? undefined : EntJobtitleEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntJobtitleToJSON(value?: EntJobtitle | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Jobtitlename': value.jobtitlename,
        'edges': EntJobtitleEdgesToJSON(value.edges),
        'id': value.id,
    };
}


