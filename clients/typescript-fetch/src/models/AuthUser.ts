/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service
 *
 * The version of the OpenAPI document: development
 * Contact: dsp-devops@broadinstitute.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AuthExtraPermissions } from './AuthExtraPermissions';
import {
    AuthExtraPermissionsFromJSON,
    AuthExtraPermissionsFromJSONTyped,
    AuthExtraPermissionsToJSON,
} from './AuthExtraPermissions';
import type { AuthFirecloudAccount } from './AuthFirecloudAccount';
import {
    AuthFirecloudAccountFromJSON,
    AuthFirecloudAccountFromJSONTyped,
    AuthFirecloudAccountToJSON,
} from './AuthFirecloudAccount';

/**
 * 
 * @export
 * @interface AuthUser
 */
export interface AuthUser {
    /**
     * 
     * @type {string}
     * @memberof AuthUser
     */
    authenticatedEmail?: string;
    /**
     * 
     * @type {AuthExtraPermissions}
     * @memberof AuthUser
     */
    matchedExtraPermissions?: AuthExtraPermissions;
    /**
     * 
     * @type {AuthFirecloudAccount}
     * @memberof AuthUser
     */
    matchedFirecloudAccount?: AuthFirecloudAccount;
}

/**
 * Check if a given object implements the AuthUser interface.
 */
export function instanceOfAuthUser(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AuthUserFromJSON(json: any): AuthUser {
    return AuthUserFromJSONTyped(json, false);
}

export function AuthUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authenticatedEmail': !exists(json, 'authenticatedEmail') ? undefined : json['authenticatedEmail'],
        'matchedExtraPermissions': !exists(json, 'matchedExtraPermissions') ? undefined : AuthExtraPermissionsFromJSON(json['matchedExtraPermissions']),
        'matchedFirecloudAccount': !exists(json, 'matchedFirecloudAccount') ? undefined : AuthFirecloudAccountFromJSON(json['matchedFirecloudAccount']),
    };
}

export function AuthUserToJSON(value?: AuthUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authenticatedEmail': value.authenticatedEmail,
        'matchedExtraPermissions': AuthExtraPermissionsToJSON(value.matchedExtraPermissions),
        'matchedFirecloudAccount': AuthFirecloudAccountToJSON(value.matchedFirecloudAccount),
    };
}

