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
/**
 * 
 * @export
 * @interface ControllersJobtitle
 */
export interface ControllersJobtitle {
    /**
     * 
     * @type {string}
     * @memberof ControllersJobtitle
     */
    jobtitlename?: string;
}

export function ControllersJobtitleFromJSON(json: any): ControllersJobtitle {
    return ControllersJobtitleFromJSONTyped(json, false);
}

export function ControllersJobtitleFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersJobtitle {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'jobtitlename': !exists(json, 'jobtitlename') ? undefined : json['jobtitlename'],
    };
}

export function ControllersJobtitleToJSON(value?: ControllersJobtitle | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'jobtitlename': value.jobtitlename,
    };
}


