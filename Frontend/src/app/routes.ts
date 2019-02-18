import {HomeComponent} from './home/home'
import {StandingsComponent} from './standings/standings'
import {TeamsComponent} from './teams/teams'

import {Routes} from "@angular/router";
import {ManageComponent} from "./manage/manage";
import {ManageLeagueComponent} from "./manage/league/manage-league";
import {ManageTeamsComponent} from "./manage/teams/manage-teams";
import {ManagePermissionsComponent} from "./manage/permissions/manage-permissions";
import {ManageDatesComponent} from "./manage/dates/manage-dates";
import {ManagePlayersComponent} from "./manage/players/manage-players";
import {ManageGamesComponent} from "./manage/games/manage-games";
import {LoginComponent} from "./login/login";
import {SignupComponent} from "./signup/signup";
import {GamesComponent} from "./games/games";
import {TournamentRegistrationComponent} from "./tournamentRegistration/tournament-registration";
import {LeaguesComponent} from "./leagues/leagues";
import {LeagueCreationComponent} from "./leagueCreation/league-creation";
import {StatsComponent} from "./stats/stats";
import {RulesComponent} from "./rules/rules";
import {ManageRulesComponent} from "./manage/rules/manage-rules";
import {ManageScheduleComponent} from "./manage/schedule/manage-schedule";

export const ELM_ROUTES: Routes = [
    {path: '', component: HomeComponent, pathMatch: 'full', data: {}},
    {path: 'rules', component: RulesComponent, data: {}},
    {path: 'standings', component: StandingsComponent, data: {}},
    {path: 'teams/:id', component: TeamsComponent, data: {}},
    {path: 'games', component: GamesComponent, data: {}},
    {path: 'login', component: LoginComponent, data: {}},
    {path: 'leagues', component: LeaguesComponent, data: {}},
    {path: 'leagueCreation', component: LeagueCreationComponent, data: {}},
    {path: 'signup', component: SignupComponent, data: {}},
    {path: 'register', component: TournamentRegistrationComponent, data: {}},
    {path: 'stats', component: StatsComponent, data: {}},
    {
        path: 'manage',
        component: ManageComponent,
        data: {},
        children: [
            {path: 'league', component: ManageLeagueComponent},
            {path: 'rules', component: ManageRulesComponent},
            {path: 'schedule', component: ManageScheduleComponent},
            {path: 'permissions', component: ManagePermissionsComponent},
            {path: 'teams', component: ManageTeamsComponent},
            {path: 'dates', component: ManageDatesComponent},
            {path: 'players', component: ManagePlayersComponent},
            {path: 'games', component: ManageGamesComponent}
        ]
    },
    {path: '**', redirectTo: ''},
];
