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
    EntPersonal,
    EntPersonalFromJSON,
    EntPersonalFromJSONTyped,
    EntPersonalToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntJobtitleEdges
 */
export interface EntJobtitleEdges {
    /**
     * Personal holds the value of the personal edge.
     * @type {Array<EntPersonal>}
     * @memberof EntJobtitleEdges
     */
    personal?: Array<EntPersonal>;
}

export function EntJobtitleEdgesFromJSON(json: any): EntJobtitleEdges {
    return EntJobtitleEdgesFromJSONTyped(json, false);
}

export function EntJobtitleEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntJobtitleEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'personal': !exists(json, 'personal') ? undefined : ((json['personal'] as Array<any>).map(EntPersonalFromJSON)),
    };
}

export function EntJobtitleEdgesToJSON(value?: EntJobtitleEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'personal': value.personal === undefined ? undefined : ((value.personal as Array<any>).map(EntPersonalToJSON)),
    };
}


