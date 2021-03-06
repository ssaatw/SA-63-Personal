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
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntJobtitle,
    EntJobtitleFromJSON,
    EntJobtitleFromJSONTyped,
    EntJobtitleToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPersonalEdges
 */
export interface EntPersonalEdges {
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntPersonalEdges
     */
    department?: EntDepartment;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPersonalEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntJobtitle}
     * @memberof EntPersonalEdges
     */
    jobtitle?: EntJobtitle;
}

export function EntPersonalEdgesFromJSON(json: any): EntPersonalEdges {
    return EntPersonalEdgesFromJSONTyped(json, false);
}

export function EntPersonalEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPersonalEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'Department') ? undefined : EntDepartmentFromJSON(json['Department']),
        'gender': !exists(json, 'Gender') ? undefined : EntGenderFromJSON(json['Gender']),
        'jobtitle': !exists(json, 'Jobtitle') ? undefined : EntJobtitleFromJSON(json['Jobtitle']),
    };
}

export function EntPersonalEdgesToJSON(value?: EntPersonalEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'department': EntDepartmentToJSON(value.department),
        'gender': EntGenderToJSON(value.gender),
        'jobtitle': EntJobtitleToJSON(value.jobtitle),
    };
}


