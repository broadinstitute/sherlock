/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service. Note: this API will try to load and return associations in responses, so clients won\'t need to make as many requests. This behavior isn\'t recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
 *
 * The version of the OpenAPI document: development
 * Contact: dsp-devops@broadinstitute.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SherlockGitCommitV3Upsert
 */
export interface SherlockGitCommitV3Upsert {
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3Upsert
     */
    committedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3Upsert
     */
    gitBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3Upsert
     */
    gitCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3Upsert
     */
    gitRepo?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockGitCommitV3Upsert
     */
    isMainBranch?: boolean;
}

/**
 * Check if a given object implements the SherlockGitCommitV3Upsert interface.
 */
export function instanceOfSherlockGitCommitV3Upsert(value: object): value is SherlockGitCommitV3Upsert {
    return true;
}

export function SherlockGitCommitV3UpsertFromJSON(json: any): SherlockGitCommitV3Upsert {
    return SherlockGitCommitV3UpsertFromJSONTyped(json, false);
}

export function SherlockGitCommitV3UpsertFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockGitCommitV3Upsert {
    if (json == null) {
        return json;
    }
    return {
        
        'committedAt': json['committedAt'] == null ? undefined : json['committedAt'],
        'gitBranch': json['gitBranch'] == null ? undefined : json['gitBranch'],
        'gitCommit': json['gitCommit'] == null ? undefined : json['gitCommit'],
        'gitRepo': json['gitRepo'] == null ? undefined : json['gitRepo'],
        'isMainBranch': json['isMainBranch'] == null ? undefined : json['isMainBranch'],
    };
}

export function SherlockGitCommitV3UpsertToJSON(value?: SherlockGitCommitV3Upsert | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'committedAt': value['committedAt'],
        'gitBranch': value['gitBranch'],
        'gitCommit': value['gitCommit'],
        'gitRepo': value['gitRepo'],
        'isMainBranch': value['isMainBranch'],
    };
}

