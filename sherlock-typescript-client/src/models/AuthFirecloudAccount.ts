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

import { exists, mapValues } from '../runtime';
import type { AuthFirecloudGroupMembership } from './AuthFirecloudGroupMembership';
import {
    AuthFirecloudGroupMembershipFromJSON,
    AuthFirecloudGroupMembershipFromJSONTyped,
    AuthFirecloudGroupMembershipToJSON,
} from './AuthFirecloudGroupMembership';

/**
 * 
 * @export
 * @interface AuthFirecloudAccount
 */
export interface AuthFirecloudAccount {
    /**
     * 
     * @type {boolean}
     * @memberof AuthFirecloudAccount
     */
    acceptedGoogleTerms?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof AuthFirecloudAccount
     */
    archived?: boolean;
    /**
     * 
     * @type {string}
     * @memberof AuthFirecloudAccount
     */
    email?: string;
    /**
     * 
     * @type {boolean}
     * @memberof AuthFirecloudAccount
     */
    enrolledIn2Fa?: boolean;
    /**
     * 
     * @type {AuthFirecloudGroupMembership}
     * @memberof AuthFirecloudAccount
     */
    groups?: AuthFirecloudGroupMembership;
    /**
     * 
     * @type {boolean}
     * @memberof AuthFirecloudAccount
     */
    suspended?: boolean;
    /**
     * 
     * @type {string}
     * @memberof AuthFirecloudAccount
     */
    suspensionReason?: string;
}

/**
 * Check if a given object implements the AuthFirecloudAccount interface.
 */
export function instanceOfAuthFirecloudAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AuthFirecloudAccountFromJSON(json: any): AuthFirecloudAccount {
    return AuthFirecloudAccountFromJSONTyped(json, false);
}

export function AuthFirecloudAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthFirecloudAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'acceptedGoogleTerms': !exists(json, 'acceptedGoogleTerms') ? undefined : json['acceptedGoogleTerms'],
        'archived': !exists(json, 'archived') ? undefined : json['archived'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'enrolledIn2Fa': !exists(json, 'enrolledIn2Fa') ? undefined : json['enrolledIn2Fa'],
        'groups': !exists(json, 'groups') ? undefined : AuthFirecloudGroupMembershipFromJSON(json['groups']),
        'suspended': !exists(json, 'suspended') ? undefined : json['suspended'],
        'suspensionReason': !exists(json, 'suspensionReason') ? undefined : json['suspensionReason'],
    };
}

export function AuthFirecloudAccountToJSON(value?: AuthFirecloudAccount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'acceptedGoogleTerms': value.acceptedGoogleTerms,
        'archived': value.archived,
        'email': value.email,
        'enrolledIn2Fa': value.enrolledIn2Fa,
        'groups': AuthFirecloudGroupMembershipToJSON(value.groups),
        'suspended': value.suspended,
        'suspensionReason': value.suspensionReason,
    };
}
